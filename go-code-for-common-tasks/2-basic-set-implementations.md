# Basic Command

**1. Display of Top Command**

show information like **tasks**, **memory**, **cpu** and **swap**. Press ‘**q**‘ to quit window.

```text
# to

```

**2. Sorting with -O \(Uppercase Letter ‘O’\).**

Press \(**Shift+O**\) to Sort field via field letter, for example press ‘**a**‘ letter to sort process with PID \(**Process ID**\).

Type any key to return to main top window with sorted **PID** order as shown in below screen. Press ‘**q**‘ to quit exit the window.

**3. Display Specific User Process**

Use top command with ‘**u**‘ option will display specific **User** process details.

```text
# top -u tecmint
```

**4. Highlight Running Process in Top**

Press ‘**z**‘ option in running top command will display running process in color which may help you to identified running process easily.

**5. Shows Absolute Path of Processes**

Press ‘**c**‘ option in running top command, it will display absolute path of running process.

**6. Change Delay or Set ‘Screen Refresh Interval’ in Top**

By default screen refresh interval is **3.0** seconds, same can be change pressing ‘**d**‘ option in running top command and change it as desired as shown below.

[  
](https://www.tecmint.com/wp-content/uploads/2012/08/Top-Set-Refresh-Time.jpg)**7. Kill running process with argument ‘k’**

You can kill a process after finding **PID** of process by pressing ‘**k**‘ option in running top command without exiting from top window as shown below.

**8. Sort by CPU Utilisation**

Press \(**Shift+P**\) to sort processes as per **CPU** utilization. See screenshot below.

**9. Renice a Process**

You can use ‘**r**‘ option to change the priority of the process also called Renice.

**10. Save Top Command Results**

To save the running top command results output to a file **/root/.toprc** use the following command.

```text
# top -n 1 -b > top-output.txt
```

**11. Getting Top Command Help**

Press ‘**h**‘ option to obtain the top command help.

**12. Exit Top Command After Specific repetition**

Top output keep refreshing until you press ‘**q**‘. With below command top command will automatically exit after 10 number of repetition.

```text
# top -n 10
```



## [Linux下查看某一进程所占用内存的方法](https://www.cnblogs.com/freeweb/p/5407105.html)

　　Linux下查看某一个进程所占用的内存，首先可以通过ps命令找到进程id，比如 ps -ef \| grep kafka 可以看到kafka这个程序的进程id

可以看到是2913，现在可以使用如下命令查看内存：

```text
top -p 2913
```

这样可以动态实时的看到CPU和内存的占用率，然后按q键回到命令行

也可直接使用ps命令查看： ps -aux \| grep kafka 

第一个标注的地方是CPU和内存占用率，后面的943100是物理内存使用量，单位是k，此时kafka大约占用943M内存

还可以查看进程的status文件： cat /proc/2913/status 

