// MinHeap Class Implementation
class MinHeap {
private:
    vector<int> heap;

    void bubbleUp(int index) {
        int parent = index / 2;
        while (index > 1 && heap[parent] > heap[index]) {
            swap(heap[parent], heap[index]);
            index = parent;
            parent = index / 2;
        }
    }

    void bubbleDown(int index) {
        int child = 2 * index;  // left child
        while (child < heap.size()) {
            // If the right child exists and is smaller than the left child, choose it.
            if (child + 1 < heap.size() && heap[child] > heap[child + 1]) {
                child++;
            }

            // If the current value is smaller than its smallest child, we're done.
            if (heap[child] >= heap[index]) {
                break;
            }

            // Swap the current value with its smallest child.
            swap(heap[child], heap[index]);
            index = child;
            child = 2 * index;  // left child
        }
    }

public:
    MinHeap() {
        heap.push_back(0);  // Dummy value to fill the 0th index
    }

    void push(int val) {
        heap.push_back(val);
        bubbleUp(heap.size() - 1);
    }

    int pop() {
        if (heap.size() <= 1) {
            return -1;
        }

        if (heap.size() == 2) {
            int root = heap.back();
            heap.pop_back();
            return root;
        }

        int root = heap[1];
        heap[1] = heap.back();
        heap.pop_back();
        bubbleDown(1);

        return root;
    }

    int top() {
        return (heap.size() > 1) ? heap[1] : -1;
    }

    void heapify(const vector<int>& arr) {
        heap.clear();
        heap.push_back(0);  // Dummy value to fill the 0th index
        heap.insert(heap.end(), arr.begin(), arr.end());
        for (int i = (heap.size() - 1) / 2; i >= 1; i--) {
            bubbleDown(i);
        }
    }
};
