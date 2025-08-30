#include <iostream>

int crop_digits(int num, int k) {
  int digits = 0;
  int temp = num;

  while (temp) {
    temp /= 10;
    digits++;
  }

  if ( k >= digits) {
    return num;
  }

  for (int i = 0; i < digits - k; i++) {
    num /= 10;
  }

  return num;
}

int main() {
  // int num = 3456789;
  int num = 3456789;
  // int num = 3456789;
  // int k = 3;
  int k = 0;
  // int k = 3;

  std::cout << "First " << k << " digits of " << num << " are: " << crop_digits(num, k) << std::endl;

  return 0;
}
