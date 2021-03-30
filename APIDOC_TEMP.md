# Directed Graph API - Temporary API Documentation

Check [README.md](README.md) for more details.

## Table of Contents
1. [Directed Graphs Endpoints](#directed-graphs-endpoints)
2. [Nodes Endpoints](#nodes-endpoints)
3. [Edges Endpoints](#edges-endpoints)


## **Directed Graphs Endpoints**


### `POST` [/v1/directed-graphs]() Create a directed graph

Request body example:
```json
{
  "name": "Graph 1"
}
```

Response Body example:
```json
{
  "data": {
    "id": 1, 
    "name": "Graph 1"
  }
}
```

Status Codes:

`201` Created


### `GET` [/v1/directed-graphs]() Lists all directed graphs

Response Body example:
```json
{
  "data": [
    {
      "id": 1, 
      "name": "Graph 1"
    }
  ]
}
```

Status Codes:

`200` OK


### `GET` [/v1/directed-graphs/*:directed-graph-id*]() Get a specific directed graphs

Response Body example:
```json
{
  "data": {
    "id": 1, 
    "name": "Graph 1"
  }
}
```

### `PUT` [/v1/directed-graphs/*:directed-graph-id*]() Update a specific directed graph

Request body example:
```json
{
  "name": "Graph 1 changed"
}
```

Status Codes:

`204` No Content


### `DELETE` [/v1/directed-graphs/*:directed-graph-id*]() Delete a specific directed graph

Status Codes:

`204` No Content



## **Nodes Endpoints**


### `POST` [/v1/directed-graphs/*:directed-graph-id*/nodes]() Create a node

Request body example:
```json
{
    "metadata": {
        "name": "Node 1",
        "any key": "Any Value"
    }
}
```

Response Body example:
```json
{
    "data": {
        "id": 1,
        "metadata": {
          "name": "Node 1",
          "any key": "Any Value"
        }
    }
}
```

Status Codes:

`201` Created


### `GET` [/v1/directed-graphs/*:directed-graph-id*/nodes]() Lists all nodes

Response Body example:
```json
{
    "data": [
        {
            "id": 1,
            "metadata": {
                "name": "Node 1",
                "any key": "Any Value"
            }
        },
        {
            "id": 2,
            "metadata": {
                "name": "Node 2",
                "any key": "Any Value",
                "any key 2": "Any Value 2"
            }
        }
    ]
}

```

Status Codes:

`200` OK


### `GET` [/v1/directed-graphs/*:directed-graph-id*/nodes/*node-id*]() Get a specific node

Response Body example:
```json
{
  "data": { 
    "id": 1,
    "metadata": {
        "name": "Node 1",
        "any key": "Any Value"
    }
  }
}
```

### `PUT` [/v1/directed-graphs/*:directed-graph-id*/nodes/*node-id*]() Update a specific node

Request body example:
```json
{
    "metadata": {
        "name": "Node 1 Changed",
        "any new key": "Any new value"
    }
}
```

Status Codes:

`204` No Content


### `DELETE` [/v1/directed-graphs/*:directed-graph-id*/nodes/*node-id*]() Delete a specific node

Status Codes:

`204` No Content



## **Edges Endpoints**


### `POST` [/v1/directed-graphs/*:directed-graph-id*/edges]() Create an edge

Request body example:
```json
{
    "tail_node_id": 1,
    "head_node_id": 2
}
```

Response Body example:
```json
{
    "data": {
        "tail_node_id": 1,
        "head_node_id": 2
    }
}
```

Status Codes:

`201` Created


### `GET` [/v1/directed-graphs/*:directed-graph-id*/edges]() Lists all edges

Response Body example:
```json
{
    "data": [
        {
            "tail_node_id": 1,
            "head_node_id": 2
        },
        {
            "tail_node_id": 3,
            "head_node_id": 4
        }
    ]
}

```

Status Codes:

`200` OK


### `GET` [/v1/directed-graphs/*:directed-graph-id*/edges/*:tail-node-id*/*:head-node-id*]() Get a specific edge

Response Body example:
```json
{
    "data": {
        "tail_node_id": 1,
        "head_node_id": 2
    }
}
```

### `DELETE` [/v1/directed-graphs/*:directed-graph-id*/edges/*:tail-node-id*/*:head-node-id*]() Delete a specific edge

Status Codes:

`204` No Content

