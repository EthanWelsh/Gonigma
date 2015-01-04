class Reflector {

public:
    Reflector(const char *reflector_file);
    char translate(char c);

private:
    char reflect[26];

};