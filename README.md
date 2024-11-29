# **Konbata**  
*A dead simple CLI tool to convert image formats.*

## Installation

To install konbata, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/FelipeMCassiano/konbata@latest
```

## **Usage**  
Konbata makes it easy to convert images between formats directly from the command line.

### **Commands and Flags**  
| **Command**            | **Flags**                                   | **Description**                                             |  
|-------------------------|---------------------------------------------|-------------------------------------------------------------|  
| `konbata [IMAGE_PATH]`  | `--png` (convert to PNG) <br> `--jpeg` (convert to JPEG) | Converts the image at the specified path to the desired format. |

---

### **Example**  
Convert an image to PNG:  
```bash
konbata image.jpg --png
```

Convert an image to JPEG:  
```bash
konbata image.png --jpeg
```
