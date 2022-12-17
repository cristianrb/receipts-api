# Description
The objective is to build a simple CRUD API to process Receipt (payment) resources. You can use either the REST API or any RPC (gRPC and so on) model. How the data is stored is up to you.

## Api Endpoints:
- Create a new item ✓
- Get list of items by ids ✓
- Get all items ✓
- Update item ✓
- Delete item ✓
- Create a new Receipt ✓
- Update a Receipt ✓
- Delete the Receipt by its id ✓
- Get the Receipt by its id ✓
- Get all Receipts ✓
- Get filtered Receipts by creation date range
- Get filtered Receipts by product name - find all Receipts containing an item with the name containing text


## Entities
### Item
- Id - a unique identifier
- ProductName - the product name as string
### Receipt
- Id - a unique identifier
- CreatedOn - the date when the given Receipt was created on
- Items: list of Receipt Items.
