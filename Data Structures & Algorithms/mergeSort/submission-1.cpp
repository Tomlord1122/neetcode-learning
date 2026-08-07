// Definition for a Pair
// class Pair {
// public:
//     int key;
//     string value;

//     Pair(int key, string value) : key(key), value(value) {}
// };
class Solution {
public:
    // Implementation of MergeSort
    vector<Pair> mergeSort(vector<Pair>& pairs) {
        mergeSortHelper(pairs, 0, pairs.size() - 1);
        return pairs;
    }

    void mergeSortHelper(vector<Pair>& pairs, int s, int e) {
        if (e - s + 1 <= 1) {
            return;
        }

        // The middle index of the array
        int m = (s + e) / 2;

        // Sort the left half
        mergeSortHelper(pairs, s, m);

        // Sort the right half
        mergeSortHelper(pairs, m + 1, e);

        // Merge sorted halves
        merge(pairs, s, m, e);
    }

    // Merge in-place
    void merge(vector<Pair>& arr, int s, int m, int e) {
        // Copy the sorted left and right halves to temp arrays
        vector<Pair> L(arr.begin() + s, arr.begin() + m + 1);
        vector<Pair> R(arr.begin() + m + 1, arr.begin() + e + 1);

        int i = 0; // index for L
        int j = 0; // index for R
        int k = s; // index for arr

        // Merge the two sorted halves into the original array
        while (i < L.size() && j < R.size()) {
            if (L[i].key <= R[j].key) {
                arr[k] = L[i];
                ++i;
            } else {
                arr[k] = R[j];
                ++j;
            }
            ++k;
        }

        // One of the halves will have elements remaining
        while (i < L.size()) {
            arr[k] = L[i];
            ++i;
            ++k;
        }
        while (j < R.size()) {
            arr[k] = R[j];
            ++j;
            ++k;
        }
    }
};
