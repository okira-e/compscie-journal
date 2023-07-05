class Solution {
    public:
        bool containsDuplicate(std::vector<int> &nums) {
            auto hashMap = std::unordered_map<int, bool>();
            for (auto num : nums) {
                if (hashMap.find(num) != hashMap.end()) {
                    return true;
                }
                hashMap[num] = true;
            }

            return false;
        }
};
