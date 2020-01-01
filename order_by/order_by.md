```
SELECT orderNumber,
orderLineNumber,
quantityOrdered * priceEach
FROM
orderdetails
ORDER BY
quantityOrdered * priceEach DESC;
```
カラム同士を計算してORDER BY指定することができる

|orderNumber|orderLineNumber|quantityOrdered * priceEach|
|----|----|----|
|10403|9|11503.14|
|10405|5|11170.52|