# FinalProject4_kel7
Link github = https://github.com/Bobby-P-dev/FinalProject3_kel7

# dokumentasi API

Domain Api = https://project4kel7.adaptable.app/


Domain Database = postgres://dkulinie:UHxJXZ5kIru0GHw-r1hDR5iWVBRSX1yE@flora.db.elephantsql.com/dkulinie

## account admin
    {
    "email":    "admin@gmail.com",
    "password": "123456"
    }

***

# User
## regist account

Method = POST

Domain = https://project4kel7.adaptable.app/users/register


request body 
```
{
    "full_name": "string",
    "email":    "string",
    "password": "string"
}
```

## Login

Method = POST

Domain = https://project4kel7.adaptable.app//users/login

request body

```
{
    "email":    "string",
    "password": "string"
}
```

## patch balance

Method = Patch

Domain = https://project4kel7.adaptable.app/users/topup

request
`
bearer token authorization
`

request body
```
{
    "balance": int
}
```

# Category
## create category

Method = POST

Domain = https://project4kel7.adaptable.app/categories/create

request
`
bearer token authorizaiton only admin
`

request body
```
  {
      "type":   "string"
  }
```

## get category

Method = GET

Domain = https://project4kel7.adaptable.app/categories/get

request
`
bearer token authorization only admin
`

## patch category

Method = PATCH

Domain = https://project4kel7.adaptable.app/categories/patch/:id

request
`
id param(int) & bearer token authorization only admin
`

request body
```
  {
    "type": "string"
  }
```

## delete category

Method = DELETE

Domain = https://project4kel7.adaptable.app/catehories/delete/:id

request
`
id param(int) & bearer token authorization only admin
`


----------------------------------------------------------------------------

# Product
## create product

Method = POST

Domain = https://project4kel7.adaptable.app/products/create

request
`
bearer token authorization only admin
`

request body
```
{
    "title":  "string",
    "price":  int,
    "stock":  int,
    "category_id":  int
}
```

## get product

Method = GET

Domain = https://project4kel7.adaptable.app/products/get

request
`
bearer token authorization
`

## put product

Method = PUT

Domain = https://project4kel7.adaptable.app/products/put/:id

request
`
id param(int) & bearer token authorization only admin
`

request body
```
    {
      "title":   "string",
      "price":   int,
      "stock":   int,
      "category_id": int
    }
```

## delete product

Method = DELETE

Domain = https://project4kel7.adaptable.app/products/delete/:id

request
`
id param(int) & bearer token authorization only admin
`

# Transaction History
## create transaction

Method = POST

Domain = https://project4kel7.adaptable.app/transactions/create

request
`
bearer token authorization
`

request body
```
{
    "product_id":  int,
    "quantity":    int
}
```

## get bill by customer

Method = GET

Domain = https://project4kel7.adaptable.app/transactions/my-transactions

request
`
bearer token authorization
`

## get bill by admin

Method = GET

Domain = https://project4kel7.adaptable.app/transactions/user-transactions/:id

request
`
bearer token authorization
`
