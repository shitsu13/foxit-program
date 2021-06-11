## Foxit - SQL



Write a SQL query command for a report that provides the following information for each document in the **Document** table, regardless if there is a content for each of those documents.

```sql
Author, Title, Metadata, Content
```

- Schema as follows

  ```sql
  Table: Document
  
  +-------------+---------+
  | Column Name | Type    |
  +-------------+---------+
  | DocumentId  | int     | PK
  | Author      | varchar |
  | Title       | varchar |
  +-------------+---------+
  
  Table: Content
  
  +-------------+---------+
  | Column Name | Type    |
  +-------------+---------+
  | ContentId   | int     | PK
  | DocumentId  | int     |
  | Metadata    | varchar |
  | Content     | varchar |
  +-------------+---------+
  ```



### Idea.

- As we know, Document & Content Table are relational model.
- Then, I choose to use `JOIN` syntax to query.
- There's one thing I notice. It's a query that's regardless if there is a content for each of those documents.
- So, I change it to use `LEFT JOIN` whether  each document has content or not. 
- And if it doesn't have the one, it will show that columns of Content are **NULL**. 
- By using `IFNULL()`, I think it maybe looks like just normal blank data and that will be fine.



### Answer.

```mariadb
select 
	d.Author as Author, d.Title as Title, IFNULL(c.Metadata, "") as Metadata, IFNULL(c.Content, "") as Content
from 
	Document as d
left join 
	Content as c on c.DocumentId = d.DocumentId
```

