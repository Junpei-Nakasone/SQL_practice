
### buildings テーブルを作成する
```
CREATE TABLE buildings (
    building_no INT PRIMARY KEY AUTO_INCREMENT,
    building_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL
);
```

### rooms テーブルを作成する
```
CREATE TABLE rooms (
    room_no INT PRIMARY KEY AUTO_INCREMENT,
    room_name VARCHAR(255) NOT NULL,
    building_no INT NOT NULL,
    FOREIGN KEY (building_no)
        REFERENCES buildings (building_no)
        ON DELETE CASCADE
);
```
FOREIGN KEY制約にON DELETE CASCADE句が付いている


### buildingsmテーブルに行を挿入する
```
INSERT INTO buildings(building_name, address)
VALUES('ACME headquaters', '3950 North 1st Street'),
      ('ACME Sales', '5000 North 1st Street');
```


### buildingsテーブルの中身を確認する
```
SELECT * FROM buildings;
```
二つの行が挿入されている

|building_no|building_name|address|
|:---|:---|:---|
|1|ACME_Headquaters|3950 North 1st Street|
|2|ACME sales|5000 North 1st Street|


### roomsテーブルに行を挿入する
```
INSERT INTO rooms(room_name,building_no)
VALUES('Amazon',1),
      ('War Room', 1),
      ('Office of CEO',1),
      ('Marketing',2),
      ('Showroom',2);
```

### roomsテーブルの中身を確認する
```
SELECT * FROM rooms;
```


### buildingsテーブルからbuilding no.2を削除する
```
DELETE FROM buildings
WHERE building_no = 2;
```

### roomsテーブルの中身を確認する
```
SELECT * FROM rooms;
```

ON DELETE CASCADE句の機能で、buildingsテーブルからbuilding.no2が削除された時に
roomsテーブルのbuilding_noが2のレコードも削除されている