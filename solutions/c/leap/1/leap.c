#include "leap.h"

bool leap_year(int year) {
    if (year % 4 != 0) {
        return false;          // Not divisible by 4
    } else if (year % 100 != 0) {
        return true;           // Divisible by 4 but not by 100
    } else if (year % 400 != 0) {
        return false;          // Divisible by 100 but not by 400
    } else {
        return true;           // Divisible by 400
    }
}
