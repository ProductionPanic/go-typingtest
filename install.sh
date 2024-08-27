#!/bin/bash
git clone https://github.com/ProductionPanic/go-typingtest.git ~/go-typingtest
cd ~/go-typingtest || exit
go build -o type-test .
chmod +x type-test
sudo mv type-test /usr/local/bin
cd ~ || exit
rm -rf ~/go-typingtest
echo "type-test installed to /usr/local/bin/type-test"
echo "run type-test to start the typing test"
echo "type-test --help for more information"

