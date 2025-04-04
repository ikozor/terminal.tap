# terminal.tap api
### Morse code commands to use in the application

## List Actions

### LIST PRODUCTS
List all the products available to buy at terminal.shop

**Example:**

command: ```.-.. .. ... - / .--. .-. --- -.. ..- -.-. - ...```

result: ```-.-. .-. --- -. --..-- / ..-. .-.. --- .-- --..-- / .- .-. - .. ... .- -. --..-- / .--..-- --- -... .--- . -.-. - / --- -... .--- . -.-. - --..--. --..-- / ... . --. ..-. .- ..- .-.. - --..-- / -.. .- .-. -.- / -- --- -.. . --..-- / ....- ----- ....-```

plaintext result: ``` CRON, FLOW, ARTISAN, [OBJECT OBJECT], SEGFAULT, DARK MODE, 404 ```  

### LIST ADDRESSES
List the addresses the user has saved

**Example:**
> Assuming the user has home and work addresses saved

command: ```.-.. .. ... - / .- -.. -.. .-. . ... ... . ...```

result: ```-.....- -. .- -- . ---... / .... --- -- . --..-- / ... - .-. . . - / .---- ---... / / .---- ..--- ...-- / .- -... -.-. / ... - --..-- / -.-. .. - -.-- ---... / / -.-. .. - -.-- --..-- / ... - .- - . ---... / ... - --..-- / --.. .. .--. -.-. --- -.. . ---... / .---- ..--- ...-- ....- ..... --..-- / -.-. --- ..- -. - .-. -.-- ---... / ..- ... --..-- / .--. .... --- -. . / -. ..- -- -... . .-. ---... / ....- ..--- ----- ..... ..... ..... -.... ----. -.... ----. .-----. --..-- / -.....- -. .- -- . ---... / .-- --- .-. -.- --..-- / ... - .-. . . - / .---- ---... / / ....- ..... -.... / .... . .-.. .-.. --- / .-.. -. --..-- / ... - .-. . . - / ..--- ---... / / ... . -.-. --- -. -.. / .--. .- .-. - --..-- / -.-. .. - -.-- ---... / / -.-. .. - -.-- --..-- / ... - .- - . ---... / ... - --..-- / --.. .. .--. -.-. --- -.. . ---... / .---- ..--- ...-- ....- ..... --..-- / -.-. --- ..- -. - .-. -.-- ---... / ..- ... --..-- / .--. .... --- -. . / -. ..- -- -... . .-. ---... / ....- ..--- ----- ..... ..... ..... -.... ----. -.... ----. .-----.```

plaintext result: ```(NAME: HOME, STREET 1:  123 ABC ST, CITY:  CITY, STATE: ST, ZIPCODE: 12345, COUNTRY: US, PHONE NUMBER: 4205556969), (NAME: WORK, STREET 1:  456 HELLO LN, STREET 2:  SECOND PART, CITY:  CITY, STATE: ST, ZIPCODE: 12345, COUNTRY: US, PHONE NUMBER: 4205556969)```

### LIST CARDS
List the cards the user has saved

**Example:**
> assuming the user has the stripe testing card 

command: ```.-.. .. ... - / -.-. .- .-. -.. ...```

result: ```-.....- ...- .. ... .- --..-- / .-.. .- ... - / ....- ---... / ....- ..--- ....- ..--- --..-- / . -..- .--. ---... / ...-- -..-. ..--- ----- ...-- ----- .-----.```

plaintext result: ```(VISA, LAST 4: 4242, EXP: 3/2030)```

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

## Address Actions

### ADDRESS ADD <ADDRESS_NAME>, <ADDRESS_STREET1>, <optional: ADDRESS_STREET2>, <ADDRESS_CITY>, <ADDRESS_STATE>, <ADDRESS_ZIPCODE>, <ADDRESS_COUNTRY>, <ADDRESS_PHONE_NUMBER>
Add new address with each part seperated by a comma (,)

**Example:** ADDRESS ADD WORK, 456 HELLO LN, SECOND PART, CITY, ST, 12345, US, 4205556969

command: ```.- -.. -.. .-. . ... ... / .- -.. -.. / .-- --- .-. -.- --..-- / ....- ..... -.... / .... . .-.. .-.. --- / .-.. -. --..-- / ... . -.-. --- -. -.. / .--. .- .-. - --..-- / -.-. .. - -.-- --..-- / ... - --..-- / .---- ..--- ...-- ....- ..... --..-- / ..- ... --..-- / ....- ..--- ----- ..... ..... ..... -.... ----. -.... ----.```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .- -.. -.. . -.. / .- -.. -.. .-. . ... ...```

plaintext result: ```SUCCESSFULLY ADDED ADDRESS```

### ADDRESS REMOVE <ADDRESS_NAME>
Remove all the addresses with that name

**Example:** ADDRESS REMOVE WORK

command: ```.- -.. -.. .-. . ... ... / .-. . -- --- ...- . / .-- --- .-. -.-```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .-. . -- --- ...- . -.. / .- -.. -.. .-. . ... ... ---... / .-- --- .-. -.-```

plaintext result: ```SUCCESSFULLY REMOVED ADDRESS: WORK```

### ADDRESS SET <ADDRESS_NAME>
Set the address that will be used for ordering and subscribing

**Example:** ADDRESS SET WORK

command ```.- -.. -.. .-. . ... ... / ... . - / .-- --- .-. -.-```

result: ```.- -.. -.. . ... ... / .-- --- .-. -.- / ... . - / ... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.--```

plaintext result: ```ADDRESS WORK SET SUCCESSFULLY```

## Card Actions

### CARD ADD 
Get a link to add a new card 

**Example**

command: ```-.-. .- .-. -.. / .- -.. -..```

result: ```--. --- / - --- / .... - - .--. ... ---... -..-. -..-. -.. . ...- .-.-.- - .-. -- .-.-.- ... .... -..-. ----- ----- ----- ----- ----- ----- ----- ----- / - --- / .- -.. -.. / - .... . / -.-. .- .-. -..```

plaintext result: ```GO TO HTTPS://DEV.TRM.SH/00000000 TO ADD THE CARD```
