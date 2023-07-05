#include <iostream>
#include <vector>


int main() {
    int n, q;

    std::cin >> n >> q;

    std::vector<int64_t> accArr(n + 1, 0);

    for (int i = 0; i < q; ++i) {
        int l, r, v;
        std::cin >> l >> r >> v;

        accArr[l] += v;

        if (r + 1 < accArr.size()) {
            accArr[r + 1] -= v;
        }
    }

    for (int i = 1; i < accArr.size(); ++i) {
        accArr[i] += accArr[i - 1];
    }

    for (int i = 1; i < accArr.size(); ++i) {
        if (i > 1) {
            std::cout << " ";
        }
        std::cout << accArr[i];
    }


    return 0;
}

