- user can upload a CSV file using an endpoint. The CSV file will contain tuples in the form of: 
date,amount,currency,type(expense | income),concept,category. These will be known as the raw data. 
- never-red will query a currency converter external API in case currency conversion is needed
- user uses an endpoint to upload individual expenses/incomes
- some of the inforamtion a user might want to get from the raw data could be: 
    - expenses/income, filtered by date (between), category (one or several), currency (one or several), amount (greater, equal, between, less)
    - time series of electricity and gas bills
    - time series of expenses, income and savings (income - expenses)
    - expenses percentiles, i. e., return top 10 expenses in the current month (or arbitrary period)