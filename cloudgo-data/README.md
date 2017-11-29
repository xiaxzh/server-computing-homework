## cloudgo-data - Go Version

### database/sql

* ab 测试 POST请求添加到数据库
    ```
    ab -n 1000 -c 100 -p parameter -T application/x-www-form-urlencoded "http://localhost:8000/service/userinfo"
    ```
    ![godbc-ab-insert](assets/images/godbc-ab-insert.png)

* ab 测试 GET请求查看所有用户
    ```
    ab -n 1000 -c 100 http://localhost:8000/service/userinfo\?userid=
    ```
    ![godbc-ab-findall](assets/images/godbc-ab-findall.png)

* ab 测试 GET请求查看某个用户
    ```
    ab -n 1000 -c 100 http://localhost:8000/service/userinfo\?userid=5
    ```
    ![godbc-ab-find5](assets/images/godbc-ab-find5.png)

### xorm

* ab 测试 POST请求添加到数据库
    ```
    ab -n 1000 -c 100 -p parameter -T application/x-www-form-urlencoded "http://localhost:8000/service/userinfo/orm"
    ```
    ![orm-ab-insert](assets/images/orm-ab-insert.png)

* ab 测试 GET请求查看所有用户
    ```
    ab -n 1000 -c 100 http://localhost:8000/service/userinfo/orm\?userid=
    ```
    ![orm-ab-findall](assets/images/orm-ab-findall.png)

* ab 测试 GET请求查看某个用户
    ```
    ab -n 1000 -c 100 http://localhost:8000/service/userinfo/orm\?userid=5
    ```
    ![orm-ab-find5](assets/images/orm-ab-find5.png)