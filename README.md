![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/bazdalaz/mkload?color=green&label=mkload&style=plastic)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/bazdalaz/mkload?style=plastic)
# mkload
Simple program to generate loading scripts for PM after failure (PS or other reason).
## Prerequests:
* Rename .env_example file to .evn
* Provide the current path to the MAIN directory on your system.
## Usage:
  ```
  mkload.exe -p 800 -n b -s
 ```
 ### Parameters & flags list:
  ```
  -n string
        LCN [a/b] (default "A")
  -p string
        the plant (no default)
  -s    create a json file with programs data
  -h 	see this help message
```
If the **-s** flag is used, the program will generate one additional JSON file with actual sequence data for the selected plant. All files are generated in **same directory** where your binary is running.

	
