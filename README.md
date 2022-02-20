# Pento Backend challenge

## Running the project

In order to run the project there should be a running Postgres database with name `payroll`.
In the `.env` file there are the config values.

Using the `go run server.go` the project can be run. To have initial data in the database you can run the sql script from the `scripts/add_data.sql`.

The following query returns a payroll summary for the specified year, month and country:
```
query {
  payrollSummary(year: 2021, month: 11, country: ITALY) {
    gross
    net
    bonus
    taxes {
      name
      value
    }
    type
    user {
      firstName
    }
  }
}
```

The following mutation saves the payroll in the database:

```
mutation {
  addPayroll(data: {userId: 1, country: ITALY, grossSalary: 5500, year: 2021, month: 8, bonus: 500})
}
```

## Project description
The company Pento wants to automate the payroll calculation for its employees.

Each employee has a yearly gross salary, and on top of that, the company can give a net monthly bonus at his discretion.

Pento is a multinational corporation based in Italy, 
and it is expanding in France, and the employee can be hired only in one of those two countries, 
to calculate the net salary each of the countries has a different process:

- Italy:
    - The employee has 14 salaries per year; August and December have an extra salary
    - The employee pays 25% of taxes
    - The employee pays 7% for the national health insurance
- France:
    - The employee has 12 salaries per year
    - The employee pays 30% of taxes
    - The employee does not pay for a national health insurance

Pento wants to have a payroll summary each month for every employee with the gross/net salary breakdown.

Keep in mind that there are two kinds of payrolls current/past payrolls named "Real Payroll" and future payrolls named "Future Preview".

Pento is a tech company and needs to have an API server to serve this purpose.

## Getting started

You can fork this repo and use the fork as a basis for your project.

In the job description, you can see what is our tech stack, 
and we would love to see the code challenge based on those technologies, 
but feel free to use any other technology you think is needed explaining your reason behind the choice you did.

Give us an implementation that has idiomatic code and has reasons for your decisions.

## Timing

Get back to us when you have a timeline for when you are done.
