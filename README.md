# Budget Tracker - Teapot 418
This test task was done by team `TEAPOT-418` for `BEST::HACKath0n`, which involves developing a budget tracker website. The backend stack for this project is Golang, and the frontend stack is Vue.js, Tailwind.css. Below, you will find all the necessary information and descriptions related to the test task.

## Description
The goal of this test task is to create a budget tracker website that allows users to track their expenses and manage their budgets. 
The website has a user-friendly interface and provide functionalities such as operating their expenses, incomes, deposits and loans, categorizing transactions and generating reports.

**Team**:
- [Roman Skok](https://github.com/romesk) - backend
- [Boryslav Ziubrytskyi](https://github.com/BoryslavGlov) - backend
- [Mykola Balii](https://github.com/Kolia913) - frontend
- [Nasobko Tymofii](https://github.com/TimofiyJ) - QA

**This project uses**:
- [GitHub Actions](https://docs.github.com/en/actions)
- [AWS RDS](https://aws.amazon.com/rds/), [MySQL](https://aws.amazon.com/rds/mysql/)
- [AWS EC2](https://aws.amazon.com/ec2/)
- [Docker](https://www.docker.com/), [docker-compose](https://docs.docker.com/compose/)
- [Golang](https://go.dev/) (and various packages)
- [Vue.js](https://vuejs.org/), [Tailwind.css](https://tailwindcss.com/), [Pinia](https://pinia.vuejs.org/)
- [JWT tokens](https://jwt.io/)

> The website is hosted on AWS and can be reached using next address: http://13.48.114.197/
> 
> The website REST API is public and can be reached using next address: http://13.48.114.197:8080/
>
> Website API Swagger docs can be found in the next repo: https://github.com/418-Teapot-Team/swagger 
## How does the project actually work?

### Sign-in & Sign-up Page 
Website includes a sign-in and sign-up page to manage user accounts. Users can create a new account by providing their desired name, email address, and password. Existing users can sign in using their credentials to access their personalized dashboard and track their budget.

### Dashboard
Upon successful sign-in, users are redirected to their personalized dashboard. The dashboard serves as the main hub to see budgets and tracking expenses. It provides an overview of the user's financial information, including the Income Net for this month, chart to monitor income nets changing each month and recent transactions.

On the page header there is located currencies for USD, EUR, PLN that are taken from https://goverla.ua usings their API: https://api.goverla.ua/graphql. 

### Expenses
The expenses section allows users to add and manage their expenses. Users can input the details of each expense, such as the amount spent and a description. They can categorize expenses into different categories, such as groceries, entertainment, cafe, etc. The expenses section provides a list or a graphical representation of expenses, helping users track their spending patterns over time.

Users can sort their expenses by creation date and amount and filter it by category.

### Incomes
The incomes section enables users to record their sources of income. Users can enter the income amount, source to keep track of their earnings. This feature allows users to have a comprehensive view of their financial inflows and outflows.

Users can sort their incomes by creation date and amount and filter it by category.

### Deposits
The deposits section allows users to track their deposits or savings. Users can input details about the deposit, such as the amount, name, date, term and interest. This feature helps users keep track of their savings and monitor their progress towards financial goals.

Users can sort their deposits by opened date and amount. Deposits earning are increased automatically each month.

### Loans
The loans section allows users to manage their loans or debts. Users can add information about their loans, such as the loan amount, name, interest rate, and repayment terms. This feature assists users in tracking their loan balances and repayment progress, ensuring they stay on top of their financial obligations.

Users can sort their loans by opened date and amount. Credits payments are increased automatically each month.

### Statistic
The statistics section provides users with visual representations and insights into their financial data. Users can view charts, graphs, and other statistical representations of their expenses and incomes. This feature helps users analyze their financial habits, identify trends, and make informed decisions to improve their financial well-being.

On the charts there is shown how user's expenses/incomes changed during the selected term, there is used average values of transactions for building the plot. On the diagram user can see impact of each category.


[//]: # (## Technical side)

