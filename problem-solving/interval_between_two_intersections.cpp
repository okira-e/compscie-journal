#include <iostream>

using std::string;

template <typename... T>
void println(T... args) {
    (std::cout << ... << args) << std::endl;
}

int main() {
    int64_t a, b, c, d;
    std::cin >> a >> b >> c >> d;

    auto result = (a % 100) * (b % 100) * (c % 100) * (d % 100) % 100;

    if (result <= 9) {
        println(0, result);
    } else {
        println(result);
    }

    
    return 0;
}
