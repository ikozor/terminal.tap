# terminal.tap api
### Morse code commands to use in the application

## Product Actions

### PRODUCT LIST
List all available products

**Example**

command: ```.--. .-. --- -.. ..- -.-. - / .-.. .. ... -```

result: ```-.-. .-. --- -. --..-- / ..-. .-.. --- .-- --..-- / .- .-. - .. ... .- -. --..-- / .--..-- --- -... .--- . -.-. - / --- -... .--- . -.-. - --..--. --..-- / ... . --. ..-. .- ..- .-.. - --..-- / -.. .- .-. -.- / -- --- -.. . --..-- / ....- ----- ....-```

plaintext result: ```CRON, FLOW, ARTISAN, [OBJECT OBJECT], SEGFAULT, DARK MODE, 404```

### PRODUCT GET <PRODUCT_NAME>
Get specific information about a product

**Example:** PRODUCT GET SEGFAULT

command: ```.--. .-. --- -.. ..- -.-. - / --. . - / ... . --. ..-. .- ..- .-.. -```

result: ```-. .- -- . ---... / ... . --. ..-. .- ..- .-.. - --..-- / ...- .- .-. .. .- -. - ... ---... / .--..-- -.....- .. -.. ---... / ----- --..-- / -. .- -- . ---... / -- . -.. .. ..- -- / .-. --- .- ... - / --..... / .---- ..--- --- --.. / --..... / .-- .... --- .-.. . / -... . .- -. ... --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. .-----. --..--. --..-- / -.. . ... -.-. .-. .. .--. - .. --- -. ---... / .- / ... .- ...- --- .-. -.-- / -.-- . - / ... .-- . . - / -... .-.. . -. -.. / -.-. .-. . .- - . -.. / ..-. .-. --- -- / .- / -. .- - ..- .-. .- .-.. / ..-. .- ..- .-.. - / .. -. / - .... . / -.-. --- ..-. ..-. . . / -.-. .... . .-. .-. -.-- / - .... .- - / -.-. .- ..- ... . ... / .. - / - --- / -.. . ...- . .-.. --- .--. / --- -. . / -... . .- -. / .. -. ... - . .- -.. / --- ..-. / - .-- --- .-.-.-```

plaintext result: ```NAME: SEGFAULT, VARIANTS: [(ID: 0, NAME: MEDIUM ROAST | 12OZ | WHOLE BEANS, PRICE: 22.00 USD)], DESCRIPTION: A SAVORY YET SWEET BLEND CREATED FROM A NATURAL FAULT IN THE COFFEE CHERRY THAT CAUSES IT TO DEVELOP ONE BEAN INSTEAD OF TWO.```

## Cart Actions

### CART GET 
Gets the users current cart

**Example**
> assuming the user has 1 segfault and 1 404 in cart

command: ```-.-. .- .-. - / --. . -```

result: ```-.....- -. .- -- . ---... / ... . --. ..-. .- ..- .-.. - --..-- / ...- .- .-. .. .- -. - ---... / ----- --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / --.- ..- .- -. - .. - -.-- ---... / .---- .-----. --..-- / -.....- -. .- -- . ---... / ....- ----- ....- --..-- / ...- .- .-. .. .- -. - ---... / ----- --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / --.- ..- .- -. - .. - -.-- ---... / .---- .-----. / - --- - .- .-.. ---... / ....- ....- .-.-.- ----- ----- / ..- ... -..```

plaintext result: ```(NAME: SEGFAULT, VARIANT: 0, PRICE: 22.00 USD, QUANTITY: 1), (NAME: 404, VARIANT: 0, PRICE: 22.00 USD, QUANTITY: 1) TOTAL: 44.00 USD```

### CART ADD <PRODUCT_NAME> <VARIANT_ID> <optional: QUANTITY>
Add product to the cart, can specify quantity, if quantity not specified, is defaulted to 1

**Example:** CART ADD SEGFAULT 0 2

command: ```-.-. .- .-. - / .- -.. -.. / ... . --. ..-. .- ..- .-.. - / ----- / ..---```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .- -.. -.. . -.. / - --- / -.-. .- .-. -```

plaintext result: ```SUCCESSFULLY ADDED TO CART```

### CART REMOVE <PRODUCT_NAME> <VARIANT_ID> <optional: QUANTITY>
Remove product from the cart, can specify quantity, if quantity is not specified, will remove the whole product from the cart

**Example:** CART REMOVE SEGFAULT 0

command: ```-.-. .- .-. - / .-. . -- --- ...- . / ... . --. ..-. .- ..- .-.. - / -----```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .-. . -- --- ...- . -.. / .. - . -- / ..-. .-. --- -- / -.-. .- .-. -```

plaintext result: ```SUCCESSFULLY REMOVED ITEM FROM CART```

### CART ORDER 
Convert the current cart to an order with set card and address

**Example**
> Assuming the user already set the card and address

command: ```-.-. .- .-. - / --- .-. -.. . .-.```

result: ```- .-. .- -.-. -.- .. -. --. / .. -. ..-. --- ---... / ... . .-. ...- .. -.-. . ---... / ..- ... .--. ... / --. .-. --- ..- -. -.. / .- -.. ...- .- -. - .- --. . --..-- / -. ..- -- -... . .-. ---... / ----. ..--- ...-- ....- -.... ----. ----- ...-- ....- --... ----- .---- -.... --... ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- .---- ----. --..-- / ..- .-. .-.. ---... / .... - - .--. ... ---... -..-. -..-. - --- --- .-.. ... .-.-.- ..- ... .--. ... .-.-.- -.-. --- -- -..-. --. --- -..-. - .-. .- -.-. -.- -.-. --- -. ..-. .. .-. -- .- -.-. - .. --- -. -...... .. -. .--. ..- - ..--.. --- .-. .. --. - .-. .- -.-. -.- -. ..- -- -...- ----. ..--- ...-- ....- -.... ----. ----- ...-- ....- --... ----- .---- -.... --... ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- .---- ----.```

plaintext result: ```TRACKING INFO: SERVICE: USPS Ground Advantage, NUMBER: 92346903470167000000000019, URL: https://tools.usps.com/go/TrackConfirmAction_input?origTrackNum=92346903470167000000000019```

## Address Actions

### ADDRESS LIST
List the addresses the user has saved

**Example:**
> Assuming the user has home and work addresses saved

command: ```.- -.. -.. .-. . ... ... / .-.. .. ... -```

result: ```-.....- -. .- -- . ---... / .... --- -- . --..-- / ... - .-. . . - / .---- ---... / / .---- ..--- ...-- / .- -... -.-. / ... - --..-- / -.-. .. - -.-- ---... / / -.-. .. - -.-- --..-- / ... - .- - . ---... / ... - --..-- / --.. .. .--. -.-. --- -.. . ---... / .---- ..--- ...-- ....- ..... --..-- / -.-. --- ..- -. - .-. -.-- ---... / ..- ... --..-- / .--. .... --- -. . / -. ..- -- -... . .-. ---... / ....- ..--- ----- ..... ..... ..... -.... ----. -.... ----. .-----. --..-- / -.....- -. .- -- . ---... / .-- --- .-. -.- --..-- / ... - .-. . . - / .---- ---... / / ....- ..... -.... / .... . .-.. .-.. --- / .-.. -. --..-- / ... - .-. . . - / ..--- ---... / / ... . -.-. --- -. -.. / .--. .- .-. - --..-- / -.-. .. - -.-- ---... / / -.-. .. - -.-- --..-- / ... - .- - . ---... / ... - --..-- / --.. .. .--. -.-. --- -.. . ---... / .---- ..--- ...-- ....- ..... --..-- / -.-. --- ..- -. - .-. -.-- ---... / ..- ... --..-- / .--. .... --- -. . / -. ..- -- -... . .-. ---... / ....- ..--- ----- ..... ..... ..... -.... ----. -.... ----. .-----.```

plaintext result: ```(NAME: HOME, STREET 1:  123 ABC ST, CITY:  CITY, STATE: ST, ZIPCODE: 12345, COUNTRY: US, PHONE NUMBER: 4205556969), (NAME: WORK, STREET 1:  456 HELLO LN, STREET 2:  SECOND PART, CITY:  CITY, STATE: ST, ZIPCODE: 12345, COUNTRY: US, PHONE NUMBER: 4205556969)```

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

### CARD LIST
List the cards the user has saved

**Example:**
> assuming the user has the stripe testing card 

command: ```-.-. .- .-. -.. / .-.. .. ... -```

result: ```-.....- ...- .. ... .- --..-- / .-.. .- ... - / ....- ---... / ....- ..--- ....- ..--- --..-- / . -..- .--. ---... / ...-- -..-. ..--- ----- ...-- ----- .-----.```

plaintext result: ```(VISA, LAST 4: 4242, EXP: 3/2030)```

### CARD ADD 
Get a link to add a new card 

**Example**

command: ```-.-. .- .-. -.. / .- -.. -..```

result: ```--. --- / - --- / .... - - .--. ... ---... -..-. -..-. -.. . ...- .-.-.- - .-. -- .-.-.- ... .... -..-. ----- ----- ----- ----- ----- ----- ----- ----- / - --- / .- -.. -.. / - .... . / -.-. .- .-. -..```

plaintext result: ```GO TO HTTPS://DEV.TRM.SH/00000000 TO ADD THE CARD```

### CARD REMOVE <CARD_LAST_4>
Remove card based on last 4 digits

**Example** CARD REMOVE 4242 

command: ```-.-. .- .-. -.. / .-. . -- --- ...- . / ....- ..--- ....- ..---```

result: ```-.-. .- .-. -.. / ....- ..--- ....- ..--- / ... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / .-. . -- --- ...- . -..```

plaintext result: ```CARD 4242 SUCCESSFULLY REMOVED```

### CARD SET <CARD_LAST_4>
Set card based on last 4 digits

**Example** CARD SET 4242 

command: ```-.-. .- .-. -.. / ... . - / ....- ..--- ....- ..---```

result: ```-.-. .- .-. -.. / ....- ..--- ....- ..--- / ... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / ... . -```

plaintext result: ```CARD 4242 SUCCESSFULLY SET```

## Profile Actions

### PROFILE GET 
Get the current users Profile

**Example**

command: ```.--. .-. --- ..-. .. .-.. . / --. . -```

result: ```-. .- -- . ---... / -... .. .-.. .-.. -.-- / -... --- -... --..-- / . -- .- .. .-.. ---... / -... .. .-.. .-.. -.-- ......- --. -- .- .. .-.. .-.-.- -.-. --- --```

plaintext result: ```NAME: BILLY BOB, EMAIL: BILLY@GMAIL.COM```

### PROFILE UPDATE NAME <PROFILE_NAME>
Update the users name 

**Example** PROFILE UPDATE NAME JOHN DOE

command: ```.--. .-. --- ..-. .. .-.. . / ..- .--. -.. .- - . / -. .- -- . / .--- --- .... -. / -.. --- .

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / ... . - / -. .- -- .```

plaintext result: ```SUCCESSFULLY SET NAME```

### PROFILE UPDATE EMAIL <PROFILE_EMAIL>

**Example** PROFILE UPDATE EMAIL JOHNDOE@GMAIL.COM

command: ```.--. .-. --- ..-. .. .-.. . / ..- .--. -.. .- - . / . -- .- .. .-.. / .--- --- .... -. -.. --- . .--.-. --. -- .- .. .-.. .-.-.- -.-. --- --```

result: ```... ..- -.-. -.-. . ... ... ..-. ..- .-.. .-.. -.-- / ... . - / . -- .- .. .-..```

plaintext result: ```SUCCESSFULLY SET EMAIL```

## Order Actions

### ORDER LIST
Get the list of the users orders

**Example**

command: ```--- .-. -.. . .-. / .-.. .. ... -```

result: ```-.....- .. -.. ---... / ----- --..-- / .- -- --- ..- -. - ---... / ..--- ..--- ----- ----- .-----.```

plaintext result: ```(ID: 0, AMOUNT: 2200)```

### ORDER GET <ORDER_ID>
Get details about the order by the id

**Example** ORDER GET 0

command: ```--- .-. -.. . .-. / --. . - / -----```

result: ```... .... .. .--. .--. .. -. --. ---... / .... --- -- . --..-- / .- -- --- ..- -. - ---... / .--..-- ... ..- -... - --- - .- .-.. ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / ... .... .. .--. .--. .. -. --. ---... / ---.. .-.-.- ----- ----- / ..- ... -.. --..--. --..-- / - .-. .- -.-. -.- .. -. --. ---... / .--..-- ... . .-. ...- .. -.-. . ---... / --..-- / -. ..- -- -... . .-. ---... / ----. ..--- ----- ----- .---- ----. ----- ...-- ....- --... ----- .---- -.... --... ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- .---- ---.. --..--. --..-- / .. - . -- ... ---... / .--..-- -.....- -. .- -- . ---... / .--..-- --- -... .--- . -.-. - / --- -... .--- . -.-. - --..--. --..-- / --.- ..- .- -. - .. - -.-- ---... / .---- --..-- / .--. .-. .. -.-. . ---... / ..--- ..--- .-.-.- ----- ----- / ..- ... -.. --..-- / ...- .- .-. .. .- -. - ---... / .-.. .. --. .... - / .-. --- .- ... - / -..--. / .---- ..--- --- --.. / -..--. / .-- .... --- .-.. . / -... . .- -. ... .-----. --..--. --..--.```

plaintext result: ```SHIPPING: HOME, AMOUNT: [SUBTOTAL: 22.00 USD, SHIPPING: 8.00 USD], TRACKING: [SERVICE: , NUMBER: 92001903470167000000000018], ITEMS: [(NAME: [OBJECT OBJECT], QUANTITY: 1, PRICE: 22.00 USD, VARIANT: LIGHT ROAST | 12OZ | WHOLE BEANS)]]```

