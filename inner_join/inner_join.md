```
SELECT 
    productCode,
    productName,
    textDescription
FROM
    products t1
INNER JOIN productlines t2
    ON t1.productline = t2.productline;
```
productsテーブルとproductlineテーブルを結合して値を取得

```
SELECT 
    productCode,
    productName,
    textDescription
FROM
    products t1
INNER JOIN productlines t2
    ON t1.productline = t2.productline
WHERE t1.productCode = "S10_1678";
```
productsテーブルとproductlineテーブルを結合してwhere句を追加して値を取得