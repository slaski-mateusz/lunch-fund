# lunch-fund
Lunch fund settlement application.

# Purpose
This is simple application allowing to do settlements in group of people often ordering food together or paying in bars and restaurants.

# Functions

## Fund members
List of peple participating in found

### Adding member
Any time new member may be added

### Removing member
Member may be removed only when her/his settlement to others is 0

### Reminder
Members having debt to other would get automatical cyclic reminders

## Orders
Order is set of actions related with ordering and paying for meal.

### Order creation
To create order one of found members should create order. Add subset of fund members to order (they would be called "Order members")

### Order payment
Someboty (usually one of order members, but may be also other found member) pays for meal and put the amount including delivery and tips to order. This person is called order founder.

### Order calculation
One of members add to order amount for each order member
Tip and delivery costs are eually divided between all order members.
Calculations are done in 1/100 as modulo and rest is put to order member account (reward for trouble)

## Found equalisation
All if founder has debt with one of members this order part is substracted.

## Returns
It is possble to return money to order founder and annotate this in the system

# Database Schema

## Table "Members"
ID Autonumber Unique
Name Text Unique
Email Text Unique
Phone Text Unique
Avatar Blob(png bitmap)
Active Bool

## Table "Orders"
ID Autonumber Unique
Timestamp UnixTimestamp
Name Text
Founder INT foregin key to Members.ID
DeliveryCost Int
TipCost Int


## Table "Orders Details"
OrderID INT foregin key to Orders.ID
MemberID INT foregin key to Members.ID


## Table "Debts"
DebtorID INT foregin key to Members.ID
CreditorID INT foregin key to Members.ID
Amount INT

## Teams
Application would support multiple teams by creating separate Sqlite file per team 


# Global management
When application is started 1st time it check if SuperUser password exist if no ask for it.
Super user password may be reset via command line by flag "--resetSuper" 
Super user may list teams not used for long time. Send warning to members and delete teams. 


# API

"/" Documentation

"/members"
GET - List all members - Can do all members
PUT - Add new member - Can do team admin
POST - Update selected member - Can do team admin
DELETE - Delete selected member (/members?id=1234) - Can do team admin

"/orders"
GET - List all orders - Can do all members
PUT - Add new order - Can do all members
POST - Update selected order - Can do creator or founder
DELETE - Delete selected order (/members?id=1234) - Can do team admin

"/debts" 
GET - Returns all actual debts - All if logged user is admin. For ordinary member only if member is debtor or creditor.

"/teams"
GET - Returns list of teams. Possible only to global SuperUser.
PUT - Add team with 1st member marked as team admin. Possible for anyone who register team.
POST - Rename selected team. Possible only to global SuperUser and team admin.
DELETE - Delete selected team. Possible only to global SuperUser and team admin.
