#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    int n;
    cin >> n;
    vector<int> numbers(n);
    for (int i = 0; i < n; ++i) {
        cin >> numbers[i];
    }
    sort(numbers.begin(), numbers.end(), greater<>());
    long long maxPair = (long long)numbers[0] * (long long)numbers[1];

    cout << maxPair << "\n";
    return 0;
}

