/*
package main
 
import "fmt"
 
type Reader interface{
    read()
}
 
type Writer interface{
    write(string)
}
 
type ReaderWriter interface{
    Reader
    Writer
}
 
func writeToStream(writer ReaderWriter, text string){
    writer.write(text)
}
func readFromStream(reader ReaderWriter){
    reader.read()
}
 
type File struct{
    text string
}
 
func (f *File) read(){
    fmt.Println(f.text)
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}
 
func main() {
     
    myFile := &File{}
    writeToStream(myFile, "hello world")
    readFromStream(myFile)
    writeToStream(myFile, "lolly bomb")
    readFromStream(myFile)
}
*/
package main
import "fmt"
  
type Vehicle interface{
    move()
}
 
type Car struct{ model string}
type Aircraft struct{ model string}
  
  
func (c Car) move(){
    fmt.Println(c.model, "едет")
}
func (a Aircraft) move(){
    fmt.Println(a.model, "летит")
}
  
func main() {
      
    tesla := Car{"Tesla"}
    volvo := Car{"Volvo"}
    boeing := Aircraft{"Boeing"}
     
    vehicles := [...]Vehicle{tesla, volvo, boeing}
    for _, vehicle := range vehicles{
        vehicle.move()
    }
}