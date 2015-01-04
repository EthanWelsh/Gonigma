class Rotor {

public:

    Rotor(const char *rotor_file);
    char translate(char c);
    void setToPosition(int p);
    void rotateOnce();

private:
    int position;
    char contacts[];

};