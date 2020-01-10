package models

import(
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"errors"
	"fmt"
)

type Task struct {
	gorm.Model
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
}

func (p *Task) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	
}

func (p *Task) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (p *Task) SaveTask(db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Create(&p).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}

func (p *Task) FindAllTasks(db *gorm.DB, uid uint) (*[]Task, error) {
	var err error
	tasks := []Task{}
	fmt.Println("---<<<HERE")
	err = db.Debug().Model(&Task{}).Where("author_id = ?", uid).Limit(100).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	if len(tasks) > 0 {
		for i, _ := range tasks {
			err := db.Debug().Model(&User{}).Where("id = ?", tasks[i].AuthorID).Take(&tasks[i].Author).Error
			if err != nil {
				return &[]Task{}, err
			}
		}
	}
	return &tasks, nil
}

func (p *Task) FindTaskByID(db *gorm.DB, pid uint64) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}

func (p *Task) UpdateATask(db *gorm.DB) (*Task, error) {

	var err error

	err = db.Debug().Model(&Task{}).Where("id = ?", p.ID).Updates(Task{Title: p.Title, Content: p.Content}).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}

func (p *Task) DeleteATask(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Task{}).Where("id = ? and author_id = ?", pid, uid).Take(&Task{}).Delete(&Task{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Task not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}