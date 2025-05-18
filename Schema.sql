create  database SalesTracking;

use SalesTracking;

CREATE TABLE coustomerdetalis (
  id int(11) NOT NULL AUTO_INCREMENT,
  CustomerID varchar(10) NOT NULL,
  CustomerName varchar(100) DEFAULT NULL,
  CustomerEmail varchar(100) DEFAULT NULL,
  CustomerAddress varchar(250) DEFAULT NULL,
  Region varchar(100) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY CustomerID_IDX (CustomerID) USING BTREE
) ;



CREATE TABLE productmaster (
  id int(11) NOT NULL AUTO_INCREMENT,
  ProductID varchar(10) NOT NULL,
  ProductName varchar(100) DEFAULT NULL,
  Category varchar(30) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY ProductID_IDX (ProductID) USING BTREE
);


CREATE TABLE sales (
  id int(11) NOT NULL AUTO_INCREMENT,
  OrderID varchar(10) NOT NULL,
  DateofSale date DEFAULT NULL,
  QuantitySold float DEFAULT NULL,
  UnitPrice float DEFAULT NULL,
  Discount float DEFAULT NULL,
  ShippingCost float DEFAULT NULL,
  PaymentMethod varchar(30) DEFAULT NULL,
  CustomerID varchar(10) DEFAULT NULL,
  ProductID varchar(10) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY DateofSale_IDX (DateofSale) USING BTREE,
  KEY CustomerID_IDX (CustomerID) USING BTREE,
  KEY ProductID_IDX (ProductID) USING BTREE
);



CREATE TABLE synctracking (
  id int(11) NOT NULL AUTO_INCREMENT,
  SyncSuccess varchar(10) DEFAULT NULL,
  StartPunch datetime DEFAULT NULL,
  EndPunch datetime DEFAULT NULL,
  PRIMARY KEY (id)
);