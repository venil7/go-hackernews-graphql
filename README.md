# Go GraphQL Implementation for Hacker News API

This is a Go-based GraphQL API implementation for accessing Hacker News data, supporting queries to fetch detailed information about items such as stories, comments, and user data.

## Schema Overview

The API exposes a GraphQL schema that allows fetching data on Hacker News items. The primary type is `Item`, which includes fields such as `id`, `type`, `text`, `url`, and `title`, along with nested items for comments and replies.

### GraphQL Schema

```graphql
type Item {
  id: Int!
  deleted: Boolean!
  type: String!
  by: String!
  time: Int!
  text: String!
  dead: Boolean!
  parent: Int!
  poll: Int!
  kids: [Int!]!
  items: [Item!]!
  url: String!
  score: Int!
  title: String!
  parts: [Int!]!
  descendants: Int!
}

type Query {
  item(id: Int!): Item!
}
```

## Query Example

Here's a sample query to fetch an item by ID, along with nested items (e.g., comments on a story):

```graphql
{
  item(id: 12345) {
    id
    text
    title
    type
    items {
      id
      text
      title
      by
    }
  }
}
```

### Example Response

A sample response for the query above might look like this:

```json
{
  "data": {
    "item": {
      "id": 12345,
      "text": "",
      "title": "Distributed file storage: MogileFS",
      "type": "story",
      "items": [
        {
          "id": 12426,
          "text": "Another example of Brad Fitzpatrick showing what a hacker armed with Perl (and occasionally C) can do. He creates tools that obsolete expensive equipment in weekends.",
          "title": "",
          "by": "staunch"
        },
        {
          "id": 12505,
          "text": "I messed around with getting this setup a couple of weeks ago. It was a long process, and I moved on to something else before I ever finished.",
          "title": "",
          "by": "mattjaynes"
        }
      ]
    }
  }
}
```

## Getting Started

1. **Clone the Repository:**

   ```bash
   git clone <repo-url>
   cd <repo-name>
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the Server:**

   ```bash
   go run main.go
   ```

4. **Test the API:**

   Once running, access the GraphQL API by visiting [http://localhost:8080/graphql](http://localhost:8080/graphql).

## Notes

- **Data Source**: The data is fetched from Hacker News, and the schema is based on the Hacker News data model.
- **GraphQL Playground**: Use a tool like GraphQL Playground or Insomnia to test queries and explore the API.

## Contributions

Feel free to open issues or submit pull requests for any improvements or bug fixes!

---

