# Personal Work Summary Of Service-Agenda

### Student Name: [黄楠绚](https://github.com/freakkid)

* Student ID: 15331117

* Enail: threequarters@qq.com

* [Commits History](https://github.com/freakkid/service-agenda/commits?author=freakkid)

* Primarily responsible for:
    
    + Backend(service) development

    + [Write test file](https://github.com/freakkid/service-agenda/blob/master/service/entities/service_test.go) for service

    + Write most of apiary.apib

    + Add .travis.yml, LICENSE and take part in the USAGE writing of README.md

    + restructure the code of client and write some test files fot client

* About my work

    * I am responsible for backend development so I have to deal with database. In the beginning, I create only two tables: user and meeting. When I considered how to do with participators of meetings and after a long time I discovered that I have to create another table as weak entity. It's so complex to deal with three tables for they depend on each other. Some logical things have to use sql statements directly using xorm and xorm could not show weak entity conveniently which makes a lot of trouble. And I choose only create one user table that have satisfied requirement of homework.

    * For restful api, I have modified the api document several times because I always discoved it not so **restful** in the process of coding. First, we think client get key by json from server and finally we think session ID have better passed by setting cookie. 

    * And the api tools, **api blueprint** could not return different response and status code according to different request. I failed in all my attempts and I found an answer on stackoverflow said **api blueprint** cannot support this kind of operation. It may be a little unfriendly to client developer.

    * It takes me 50% time to design and modify the api document and make my code adapt to it. As for me, I learn some programming thought and "science popularization" about web from this course and learn coding about web development partly by myself. I get knowledge and joy on when coding. However, I could not get deep understanding of golang or other thing about web like some "dalao" of our class. :) I am so ashamed and uneasy every time I read his code and readme. (qaq)

    * What's more, I think it will be better if every developer can express himselp clearly and logically and not shirk responsibility.