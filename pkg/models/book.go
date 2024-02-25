package models
import{
	"github.com/jinzhu/gorm"
	"github.com/KP-11092000/BOOKSTORE/pkg/config"
}
var db *gorm.db
type Book struct{
	gorm.model
	Name string `gorm:""json:"name`
	Author string `json:"author"`
	Publication string `json:"publication"`

}
func init()  {
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookById(Id int64)(*Book,*gnorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}
func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}