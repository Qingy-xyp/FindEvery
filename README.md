# FindEvery
红队突破内网边界后，可能需要迅速寻找第二台主机权限来进行权限维持，那么最稳的方法就是对现有主机进行敏感信息查找，由于手动查找太慢或遗漏，故此诞生此工具
```
FindEvery -t .txt,.text,.ini,.yaml,.yml,.php,.jsp,.java,.xml,.sql,.properties -f "jdbc:" -p /Users/admin/Desktop/
```
-t 指定需要查找的文件类型，如：.txt,.text,.ini,.yaml,.yml,.php,.jsp,.java,.xml,.sql,.properties

-f 指定要查找的字段(不支持多个字段)，如：password=,jdbc:,user=,key=,ssh-,ldap:,mysqli_connect,sk-

-d 指定要查找的路径，如：/Users/admin/Desktop/

-o 指定要保存的路径

控制台会显示简略信息，详细信息请查看findout.txt
