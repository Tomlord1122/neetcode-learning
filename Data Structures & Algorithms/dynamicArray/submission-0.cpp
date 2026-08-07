#include <iostream>



using namespace std;

class DynamicArray {

private:
    int *arr;
    int length;
    int capacity;


public:

    DynamicArray(int capacity):capacity(capacity), length(0){
        arr = new int[capacity];
    }

    int get(int i) {
        return arr[i];
    }

    void set(int i, int n) {
        arr[i] = n;
    }

    void pushback(int n) {
        if (length >= capacity){
            resize();
        }
        arr[length] =n ;
        length++;
    }

    int popback() {
        if (length > 0){
            int lastElement = arr[length-1];
            length--;
            return lastElement;
        }
        return -1;
    }

    void resize() {
       int *newArr = new int[2*capacity];
        for(int i = 0; i < length; i++){
            newArr[i] = arr[i];
        }
        delete[] arr;

        arr = newArr;
        capacity *= 2;
    }

    int getSize() {
        return length;
    }

    int getCapacity() { 
        return capacity;
    }
};
