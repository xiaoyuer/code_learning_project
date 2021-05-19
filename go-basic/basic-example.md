---
description: decimal to int
---

# Basic-Example



## Golang Example - Convert\(cast\) Binary to Decimal or Decimal to Binary Number

In these posts, You will learn two programs in Go language  
The first program is to convert Binary Number to Decimal Number  
Second program to convert Decimal Number to Binary Number.

A binary number is a number based on base 2. It contains either zero or one digits. Each digit is called a bit, Example binary numbers are 1011

Decimal Number are numbers which contain numbers based on base 10

To understand this example, You should have following features in Go language.

* [Golang For Loop control structure](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-for-loop)
* [Golang For Loop Range form](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-foreach-loop)
* [Golang Operator Guide](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-operators-guide)
* [Golang Datatype guide](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-data-types-basic)
* [Beginner Guide to Functions with examples](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-beginner-guide)

#### Example Program to convert Binary to Decimal Number  <a id="example-program-to-convert-binary-to-decimal-number"></a>

There are 2 ways to convert binary number to decimal number

* Using strconv ParseInt function
* Write conversion logic manually without using inbuilt functions

The below programs takes input binary number from a user console, store it in variable binary.

**strconv ParseInt function example to Cast Binary To decimal** 

strconv package provides ParseInt function used to convert binary to decimal numbers Here is an example program

```text
package main  
  
import (  
 "fmt"  
 "strconv"  
)  
  
func main() {  
 var binary string  
 fmt.Print("Enter Binary Number:")  
 fmt.Scanln(&binary)  
 output, err := strconv.ParseInt(binary, 2, 64)  
 if err != nil {  
  fmt.Println(err)  
  return  
 }  
  
 fmt.Printf("Output %d", output)  
  
}  

```

Output is

```text
Enter Binary Number:1111  
Output 15  

```

**Manual Custom function to Cast Binary To decimal using for loop**

Created an own function, inside a function, used for loop  with Modulus, Division operator and Math pow function  
Here is an example program

```text
package main  
  
import (  
 "fmt"  
 "math"  
)  
  
func convertBinaryToDecimal(number int) int {  
 decimal := 0  
 counter := 0.0  
 remainder := 0  
  
 for number != 0 {  
  remainder = number % 10  
  decimal += remainder * int(math.Pow(2.0, counter))  
  number = number / 10  
  counter++  
 }  
 return decimal  
}  
  
func main() {  
 var binary int  
 fmt.Print("Enter Binary Number:")  
 fmt.Scanln(&binary)  
  
 fmt.Printf("Output %d", convertBinaryToDecimal(binary))  
  
}  

```

Output is

```text
Enter Binary Number:1001  
Output 9  

```

#### Example Program to convert Decimal to Binary Number  <a id="example-program-to-convert-decimal-to-binary-number"></a>

We have _2 ways to convert Decimal to Binary number in Golang_. Following are two ways

* using strconv FormatInt function 
* Manually conversion without using Inbuilt function 

Below program takes decimal input number from user console, stores it in variable decimal

**strconv FormatInt function example to cast Decimal to Binary** 

Inbuilt Standard package strconv provides FormatInt function used to convert Decimal to Binary number.  
Here is an example program

```text
package main  
  
import (  
 "fmt"  
 "strconv"  
)  
func main() {  
 var decimal int64  
 fmt.Print("Enter Decimal Number:")  
 fmt.Scanln(&decimal)  
 output := strconv.FormatInt(decimal, 2)  
 fmt.Print("Output ", output)  
  
}  

```

Output is

```text
Enter Decimal Number:15  
Output 1111  

```

**Cast Decimal to binary for loop without inbuilt function**

The following program uses Golang features For loop, Modulus and division operators.  
Here is an example program to do the conversion.

```text
package main  
  
import (  
 "fmt"  
)  
  
func convertDecimalToBinary(number int) int {  
 binary := 0  
 counter := 1  
 remainder := 0  
  
 for number != 0 {  
  remainder = number % 2  
  number = number / 2  
  binary += remainder * counter  
  counter *= 10  
  
 }  
 return binary  
}  
  
func main() {  
 var decimal int  
 fmt.Print("Enter Decimal Number:")  
 fmt.Scanln(&decimal)  
  
 fmt.Printf("Output %d", convertDecimalToBinary(decimal))  
  
}  

```

Output is

```text
Enter Decimal Number:15  
Output 1111  
```

