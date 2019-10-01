# todos los dias mi amogo!

## models


```
Todo
- id
- title
- description
- status
  - 0 created
  - 1 completed
  - 2 canceled
  - 3 deleted
- createdAt
- updatedAt
- deletedAt

Project
- id
- []Todo
- createdAt
- updatedAt
- deletedAt

Label
- id
- owner_id -> Todo, Project
- name
- createdAt
- updatedAt
- deletedAt

User
- id
- name
```
