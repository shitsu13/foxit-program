## Foxit - Go Scraper



Design and implement the crawler by any framework or packages to get data from Dcard.

leverage Dcard API and run the crawler within container.

- Scrape posts of forum with Apple and find a storage to keep datas based on content.

- Any database (SQL/NoSQL), Cloud service would be acceptable.

  ```sql
  Title, Content, Categories, Topics, MediaURL, CreatedAt
  ```

- Write log to output file.

- Scrape its datas each six hours, and check it's already scraped the origin data.

- Design backfill mechanism when posts were updated. (Optional)



### Idea.

- First, I think about how to implement and what tools  I need.

  - Dcard V2 API
  - Crawler frameworks or packages
  - Cron Job
  - Design backfill mechanism

- Try to leverage **Dcard V2 API** and use it.

  ```apl
  https://www.dcard.tw/service/api/v2/forums/{forums_alias}/posts?popular=false
  ```

  ```apl
  https://www.dcard.tw/service/api/v2/posts/{post_id}
  ```

- Watch its reponse data structure and know which columns I will use.

- Think about how to store its data and choose to use **MongoDB**.

- Then, I choose to use **Colly**. It's a powerful crawler framework.

  - Scrape any kind of web page and reponse with HTML DOM or JSON data structure.  
  - Support listeners on tracing request so that I can keep logs.
  - Sync / Async / Parallel scraping. Support concurrency per domain.
  - Fast (>1k request/s on a single core)
  - Clean API

- To keep a log file, I use **Logrus** package. For cron job scheduler, I use **Gocron** package.

- Also, I design backfill mechanism by checking its lastest post id.

  - If its latest post id is been in my database, it will mean that I've already store it.
  - If not, I will remove all posts already been there then update the new ones.



### Note

- Dcard V2 API, Forum api with posts
  - Response of  for 30 records / times by default.
  - Use **popular** as parameter to request for the latest or popular posts.
- So, I just use default as a base line for development.
- Design backfill mechanism to refresh each latest 30 records.
- Set cron job for each 6 hours, only full hours starting on the hour count towards your hours worked.
- Output log file as **scraper_log.log**.



### Prerequisite

- Any Environment to place its execution.

- Install MongoDB by using Docker or just in your host.

  - Set environment variables and refer to your environment.

    Example as below

    ```yaml
    version: "4.4"
    services:
    mongodb:
    image : mongo
    container_name: mongodb
    environment:
    - MONGO_ACCOUNT=$account
    - MONGO_PASSWORD=$password
    - MONGO_ADDR=$addr
    volumes:
    - /home/$user/mongodb/database:/data/db
    ports:
    - 27017:27017
    restart: unless-stopped
    ```

    

### How to start

- Run command as follows

```sh
go run ./src/main.go`
go build -o {app_name} ./src
./{app_name}
```


- You're ready to go.
- And you'll see it running on system console.
- Also, You can use **Ctrl+C** or **Any Terminated CMD** to shutdown it down.



### Test & Result

```json
/* 1 */
{
    "_id" : ObjectId("60c2165f609c7e087a5fd1f3"),
    "pid" : NumberLong(236221943),
    "title" : "有人注意到嗎？",
    "content" : "今天心血來潮把iPad的App進行更新\n發現了一個特殊的說明\nhttps://i.imgur.com/ZPO37Eo.jpg\nhttps://i.imgur.com/vm0leWG.jpg\n這個…是告訴我們Google 工程師也喜歡貓嗎🤣",
    "categories" : [],
    "topics" : [ 
        "Apple", 
        "iPad"
    ],
    "media_url" : [ 
        "https://i.imgur.com/ZPO37Eo.jpg", 
        "https://i.imgur.com/vm0leWG.jpg"
    ],
    "created_at" : ISODate("2021-06-10T13:33:24.901Z")
}

/* 2 */
{
    "_id" : ObjectId("60c2165f609c7e087a5fd1f4"),
    "pid" : NumberLong(236221923),
    "title" : "桃園a8德誼門市",
    "content" : "請問一下有人直接在裡面當場購買手機嗎？\n基本上他們會不會備貨？\n因為看在官網訂只有台北101能選擇而已\n想說直接到a8那間購買，有人有經驗嗎？\n\n感謝你們~~",
    "categories" : [],
    "topics" : [ 
        "iPhone", 
        "Apple"
    ],
    "media_url" : [],
    "created_at" : ISODate("2021-06-10T13:31:26.721Z")
}

/* 3 */
{
    "_id" : ObjectId("60c2165f609c7e087a5fd1f9"),
    "pid" : NumberLong(236221618),
    "title" : "iPad😩",
    "content" : "到底要買11or12.9啊\n疫情無法去店面看看\n12.9有比m1 air 2020重嗎（單手拿m1覺得ok\n不考慮實用度前提\n12.9會太大攜帶不便嗎\n\n只考慮螢幕大小and重量",
    "categories" : [],
    "topics" : [ 
        "iPad"
    ],
    "media_url" : [],
    "created_at" : ISODate("2021-06-10T12:57:24.127Z")
}
```



### Question



> Response Error, HTTP Statuscode 403.

I can't get any data by several times, but surf by browser. Even though, I mock as http client to request. It's always getting error. As I see its error, it's an error because Dcard is using **Cloudflare** as CDN Authentication. It will enter in by rendering Javascript.

![403錯誤](https://res.cloudinary.com/jiablog/dcard_api_v2/403error.png)



> Solution

- Use [Selenium](https://www.selenium.dev) or [Pyppeteer](https://pyppeteer.github.io/pyppeteer) to simulate as browser. 

  It's not fitted with such a crawler.

- Use Library such as [CloudScraper](https://github.com/venomous/cloudscraper) to bypass **Cloudflare**. It's built in requests.

  It's not supported by Golang.

  

### Reference



Colly - Fast and Elegant Scraping Framework for Gophers 

http://go-colly.org/

爬蟲 Dcard API 2.0 版本 

https://blog.jiatool.com/posts/dcard_api_v2/