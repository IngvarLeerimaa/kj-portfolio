# GraphQL Profile Page

## Objective
The objective of this project is to learn the GraphQL query language by creating your own profile page.

### Goals

- Display at least three pieces of information:
  - Basic user information such as username, name, email, and total XP.
  - Test data summary showing the number of tests taken, passed, and overall percentage of tests passed.
  - A list of all the subjects passed, showing the name, date when passed, and XP gained.

- Create at least two statistical graphs using SVG:
  - A pie chart showing the percentage of tests passed during different piscines.
  - A bar chart showing all the subjects passed; while hovering, you will see the name, date when passed, and XP gained.

- Create a login page and a logout button:
  - Login with wrong credentials is not allowed.
  - Logout button is available on the profile page.
  - Error message in red will be shown if wrong credentials are entered.

- Use normal, nested, and argument queries:
  - Normal querying: Fetches fields directly from the "user" type without any nesting.
  - Nested querying: Fetches fields from nested types like "progresses" and "transactions" within the "user" type.
  - Using arguments: Utilizes arguments in the "transactions" field to filter data based on the "type" attribute and to order the results by the "createdAt" attribute.

## Audit
[The audit questions can be found here](https://github.com/01-edu/public/tree/master/subjects/graphql/audit)

## Local Audit
To audit the project locally, create a server using the Live Server extension in Visual Studio Code or your preferred method. Navigate to the index.html file in the browser. Use your 01.edu credentials to log in and view your profile.

Your data will not be stored.

## Hosting
[The project is hosted on Netlify](https://this-is-my-graphql.netlify.app)
