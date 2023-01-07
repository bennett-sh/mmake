
#include <iostream>
#include "args.h"

void printArgs(char* argv[]) {
  for(int i = 0; i < sizeof(argv) / sizeof(char**); i++) {
    std::cout << "arg " << i << ": " << argv[i] << std::endl;
  }
}
