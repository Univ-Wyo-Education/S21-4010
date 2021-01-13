## index.js


From index.js Before:

```
    "0xe7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee": {
        "Addr": "e7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee",
        "Value": [
            { "BlockIndex": 0, "TxOffset": 8, "TxOutputPos": 0 },
            { "BlockIndex": 2, "TxOffset": 0, "TxOutputPos": 1 }
        ]
    },
    "0x9d41e5938767466af28865e1c33071f1561d57a8": {
        "Addr": "9d41e5938767466af28865e1c33071f1561d57a8",
        "Value": [
            { "BlockIndex": 8, "TxOffset": 0, "TxOutputPos": 0 }
        ]
    }
```

After:

```
    "0xe7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee": {
        "Addr": "e7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee",
        "Value": [
            { "BlockIndex": 8, "TxOffset": 0, "TxOutputPos": 1 }
        ]
    },
    "0x9d41e5938767466af28865e1c33071f1561d57a8": {
        "Addr": "9d41e5938767466af28865e1c33071f1561d57a8",
        "Value": [
            { "BlockIndex": 0, "TxOffset": 0, "TxOutputPos": 0 },
            { "BlockIndex": 8, "TxOffset": 0, "TxOutputPos": 0 }
        ]
    }
```

