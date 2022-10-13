# brow

Get Browser History in terminal.

## Usage

```
  -a    See All Available Options that can be used to query the Subsystem.
  -ascending
        Get Results in ascending order. Default is Descending Order
  -c    Copy the Browser Database. Ignored if path is provided. Use this Flag in case Browser is Open as it locks the database.
  -f string
        Specify the Field Separator which will be used in the Output. (default "|")
  -n int
        Specify Number of Results to be Displayed (default -1)
  -p string
        Specify Custom Path for History (default "/home/aditya/.config/BraveSoftware/Brave-Browser/Default/History")
  -q value
        Query for the current Subsystem (default [*])
  -s string
        Specify Subsystem. Either history(urls)/downloads (default "urls")
```

- Tested on Brave (Should Work for Most of the Chromium based Browsers)
- Check the [scripts](https://github.com/aditya-K2/brow/blob/master/scripts/examples) for examples to use it in a script

## Examples

### dmenu

https://user-images.githubusercontent.com/51816057/195517488-6db734ca-7943-47c6-aac5-f0315c7f765f.mp4

### fzf

https://user-images.githubusercontent.com/51816057/195517545-33186539-c55a-4134-8d95-71ab0497b9af.mp4





