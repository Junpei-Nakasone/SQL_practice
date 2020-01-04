```
SELECT
lastname,
firstname,
jobTitle
FROM 
employees
WHERE 
jobTitle = 'Sales Rep';
```
employeesテーブルからjotitleがSales Repのレコードのlastname, firstname, jobtitleを取得


```
SELECT
lastname,
firstname,
officeCode
FROM 
employees
WHERE 
officeCode BETWEEN 1 AND 3
ORDER BY officeCode;
```
officeCodeが1から3のレコードを取得

```
SELECT
lastname,
firstname,
officeCode
FROM 
employees
WHERE 
lastName LIKE '%son'
ORDER BY firstName;
```
lastnameがsonで終わるemployeesを取得

```
SELECT
lastname,
firstname,
officeCode
FROM 
employees
WHERE 
officeCode IN (1, 3)
ORDER BY officeCode;
```
officeCodeが1と3のemployeeを取得