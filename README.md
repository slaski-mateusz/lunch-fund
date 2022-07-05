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


