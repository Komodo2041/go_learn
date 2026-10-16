package raindrops
 import (
 
	"strconv"
)

func Convert(number int) string {
    name := ""
    if (number % 3 == 0) {
        name = name + "Pling"
    }
    if (number % 5 == 0) {
         name = name + "Plang"
    }   
    if (number % 7 == 0) {
       name = name + "Plong"
    }      
    if (name == "") {
         return strconv.Itoa(number)
    } else {
         return name;
    }
    
}
