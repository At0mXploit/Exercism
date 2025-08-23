#include "grains.h"

// Returns the number of grains on a given square (1-64)
uint64_t square(uint8_t index) {
    if (index == 0 || index > 64) {
        return 0; // Invalid square
    }
    return (uint64_t)1 << (index - 1); // 2^(index-1)
}

// Returns the total number of grains on the chessboard
uint64_t total(void) {
    uint64_t sum = 0;
    for (uint8_t i = 1; i <= 64; i++) {
        sum += square(i);
    }
    return sum;
}
