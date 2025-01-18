#!/bin/bash

if ! command -v go &> /dev/null
then 
    echo "Go is not installed. Please install Go and try again."
    exit
fi

if [ ! -f go.mod ]; then
    echo "Please run this script from the project root directory."
    exit
fi

echo "Setting up the project..."
go mod tidy

echo "Project setup completed."