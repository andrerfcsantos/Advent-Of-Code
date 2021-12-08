lines = open(r"..\inputs\2021_08.txt").readlines()


def process_line(line):
    parts = line.strip().split(" | ")
    return (parts[0].split(), parts[1].split())


displays = list(map(process_line, lines))

real_digits_to_number = {
    "abcefg": "0",
    "cf": "1",
    "acdeg": "2",
    "acdfg": "3",
    "bcdf": "4",
    "abdfg": "5",
    "abdefg": "6",
    "acf": "7",
    "abcdefg": "8",
    "abcdfg": "9",
}


def add_input_to_fingerprint(fingerprints, input, number):
    for digit in input:
        if digit not in fingerprints:
            fingerprints[digit] = [number]
        else:
            fingerprints[digit].append(number)


total = 0
for display in displays:
    inputs = display[0]
    outputs = display[1]

    input_to_number = {}
    number_to_input = {}

    encrypted_digits_to_real = {}
    encrypted_base_fingerprints = {}

    for input in inputs:
        if len(input) == 2:
            input_to_number[input] = 1
            number_to_input[1] = input
            add_input_to_fingerprint(encrypted_base_fingerprints, input, 1)
        elif len(input) == 3:
            input_to_number[input] = 7
            number_to_input[7] = input
            add_input_to_fingerprint(encrypted_base_fingerprints, input, 7)
        elif len(input) == 4:
            input_to_number[input] = 4
            number_to_input[4] = input
            add_input_to_fingerprint(encrypted_base_fingerprints, input, 4)
        elif len(input) == 7:
            input_to_number[input] = 8
            number_to_input[8] = input
            add_input_to_fingerprint(encrypted_base_fingerprints, input, 8)

    a = None
    for k, v in encrypted_base_fingerprints.items():
        if len(v) == 2 and 7 in v and 8 in v:
            encrypted_digits_to_real[k] = "a"
            a = k
            break

    digit_frequencies = {}
    for input in inputs:
        for digit in input:
            if digit not in digit_frequencies:
                digit_frequencies[digit] = 1
            else:
                digit_frequencies[digit] += 1

    for d, f in digit_frequencies.items():
        if f == 8 and d != a:
            encrypted_digits_to_real[d] = "c"
        elif f == 6:
            encrypted_digits_to_real[d] = "b"
        elif f == 4:
            encrypted_digits_to_real[d] = "e"
        elif f == 9:
            encrypted_digits_to_real[d] = "f"

    digits_left = list(set("abcdefg") - set(encrypted_digits_to_real.keys()))

    if digits_left[0] in number_to_input[4]:
        encrypted_digits_to_real[digits_left[0]] = "d"
        encrypted_digits_to_real[digits_left[1]] = "g"
    else:
        encrypted_digits_to_real[digits_left[1]] = "d"
        encrypted_digits_to_real[digits_left[0]] = "g"

    number_s = ""
    for output in outputs:
        real_output = ""

        for digit in output:
            real_output += encrypted_digits_to_real[digit]

        number_s += real_digits_to_number["".join(sorted(real_output))]

    number = int(number_s)
    total += number

print(total)
