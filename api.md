# terminal.tap api
### Morse code commands to use in the application

## List Actions

### LIST PRODUCTS
List all the products available to buy at terminal.shop

**Example:**

command: ```.-.. .. ... - / .--. .-. --- -.. ..- -.-. - ...```

result: ```-.-. .-. --- -. --..-- / ..-. .-.. --- .-- --..-- / .- .-. - .. ... .- -. --..-- / .--..-- --- -... .--- . -.-. - / --- -... .--- . -.-. - --..--. --..-- / ... . --. ..-. .- ..- .-.. - --..-- / -.. .- .-. -.- / -- --- -.. . --..-- / ....- ----- ....-```

plaintext result: ``` CRON, FLOW, ARTISAN, [OBJECT OBJECT], SEGFAULT, DARK MODE, 404 ```  

## Get Actions

### GET PRODUCT <PRODUCT_NAME>
Get specific information about a product

**Example:** GET PRODUCT SEGFAULT

command: ```--. . - / .--. .-. --- -.. ..- -.-. - / ... . --. ..-. .- ..- .-.. -```

result: ```--. . - / .--. .-. --- -.. ..- -.-. - / ... . --. ..-. .- ..- .-.. -
-. .- -- . ---... / ... . --. ..-. .- ..- .-.. - --..-- / - -.-- .--. . ---... / -- . -.. .. ..- -- / .-. --- .- ... - / -..--. / .---- ..--- --- --.. / -..--. / .-- .... --- .-.. . / -... . .- -. ... --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / -.. . ... -.-. .-. .. .--. - .. --- -. ---... / .- / ... .- ...- --- .-. -.-- / -.-- . - / ... .-- . . - / -... .-.. . -. -.. / -.-. .-. . .- - . -.. / ..-. .-. --- -- / .- / -. .- - ..- .-. .- .-.. / ..-. .- ..- .-.. - / .. -. / - .... . / -.-. --- ..-. ..-. . . / -.-. .... . .-. .-. -.-- / - .... .- - / -.-. .- ..- ... . ... / .. - / - --- / -.. . ...- . .-.. --- .--. / --- -. . / -... . .- -. / .. -. ... - . .- -.. / --- ..-. / - .-- --- .-.-.-```

plaintext result: ```NAME: SEGFAULT, TYPE: MEDIUM ROAST | 12OZ | WHOLE BEANS, PRICE: 22.00 USD, DESCRIPTION: A SAVORY YET SWEET BLEND CREATED FROM A NATURAL FAULT IN THE COFFEE CHERRY THAT CAUSES IT TO DEVELOP ONE BEAN INSTEAD OF TWO.```

### GET CART
gets the users current cart

**Example:** 
> assuming the user has 1 segfault and 1 404 in cart

command: ```--. . - / -.-. .- .-. -```

result: ```.....- -. .- -- . ---... / ... . --. ..-. .- ..- .-.. - --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / --.- ..- .- -. - .. - -.-- ---... / .---- .-----. --..-- / -.....- -. .- -- . ---... / ....- ----- ....- --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / --.- ..- .- -. - .. - -.-- ---... / .---- .-----. / - --- - .- .-.. ---... / ....- ....- .-.-.- ----- ----- / ..- ... -..```

plaintext result: ```(NAME: SEGFAULT, PRICE: 22.00 USD, QUANTITY: 1), (NAME: 404, PRICE: 22.00 USD, QUANTITY: 1) TOTAL: 44.00 USD```

## Cart Actions

### CART ADD <PRODUCT_NAME> <optional: QUANTITY>
Add product to the cart, can specify quantity, if quantity not specified, is defaulted to 1

**Example:** CART ADD SEGFAULT 2

command: ```-.-. .- .-. - / .- -.. -.. / ... . --. ..-. .- ..- .-.. - / ..---```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .- -.. -.. . -.. / - --- / -.-. .- .-. -```

plaintext result: ```SUCCESSFULLY ADDED TO CART```

### CART REMOVE <PRODUCT_NAME> <optional: QUANTITY>
Remove product from the cart, can specify quantity, if quantity is not specified, will remove the whole product from the cart

**Example:** CART REMOVE SEGFAULT

command: ```-.-. .- .-. - / .-. . -- --- ...- . / ... . --. ..-. .- ..- .-.. -```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .-. . -- --- ...- . -.. / .. - . -- / ..-. .-. --- -- / -.-. .- .-. -```

plaintext result: ```SUCCESSFULLY REMOVED ITEM FROM CART```
