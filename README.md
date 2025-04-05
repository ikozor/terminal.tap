# terminal.tap
terminal.tap allows you to order your favorite coffee from terminal.shop using morse code. The program is a REPL that reads a morse code string in the format `.--. .-. --- -.. ..- -.-. - / .-.. .. ... -`, evaluates the string, executes a command, and returns the result of the command as a morse code string. 

Go to [api docs](https://github.com/ikozor/terminal.tap/blob/main/api.md) to see all the commands you can use.

To to [morse code](https://github.com/ikozor/terminal.tap/blob/main/morsecode.md) to see the Morse code translation (I had to add some characters to work with terminal.shop api)

## Environment
To see how to set up the .env file, look at [example.env](https://github.com/ikozor/terminal.tap/blob/main/example.env).

**terminal.shop**

The .env file should contain your api key as `TERMINAL_TOKEN` and url as `TERMINAL_URL` .

To setup an api token, see [terminal.shop/api](https://www.terminal.shop/api#authentication)

***Optional:* input/output**

You can also set input and output files in the .env; if you do not set them, they will default to stdin and stdout. Set input as `INPUT` and output as `OUTPUT`. The program will only read from the input file when a change has been detected and will only read the last line.

With this, you can have another program that can read a user sending Morse code (by tapping, sound, etc,) and convert it to a Morse code string and write it to the input file. terminal.tap will read the input, process it, and print the response into the output file as a Morse code string. Then the program can read the output file and return the Morse code to the user by flashing, sound, etc.

## Use case
Morse code has been a long-neglected form of communication. Many even consider it dead. But what if you are a 19th-century train company and you need to restock your carriage with coffee, and you can only send messages with Morse code? What if your boss prefers inferior coffee and has banned ordering any other coffee, so the only way you can secretly order terminal.shop coffee is by using Morse code? Hence, terminal.tap was created for such use cases.
