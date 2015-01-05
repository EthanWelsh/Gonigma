#import "Rotor.h"
#import "Reflector.h"
#import "Plugboard.h"

class Machine {

public:

    Machine(const char *rotor1_file, const char *rotor2_file, const char *rotor3_file, const char *reflector_file, const char *plugboard_file);
    char translate(char c);
    void setRotorToPosition(int rotorNum, int position);

private:
    Rotor rotors[3];
    Reflector reflector;
    Plugboard plugboard;
};