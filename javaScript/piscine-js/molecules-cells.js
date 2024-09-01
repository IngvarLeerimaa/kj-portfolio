//RNA: that converts a DNA strand into its compliment RNA strand.
function RNA (str){
    var rna = '';
    for (var i = 0; i < str.length; i++) {
        if (str[i] === 'G') {
        rna += 'C';
        } else if (str[i] === 'C') {
        rna += 'G';
        } else if (str[i] === 'T') {
        rna += 'A';
        } else if (str[i] === 'A') {
        rna += 'U';
        }
    }
    return rna;

}

//DNA: that converts an RNA strand into its compliment DNA strand.
function DNA (str){
    var dna = '';
    for (var i = 0; i < str.length; i++) {
        if (str[i] === 'C') {
        dna += 'G';
        } else if (str[i] === 'G') {
        dna += 'C';
        } else if (str[i] === 'A') {
        dna += 'T';
        } else if (str[i] === 'U') {
        dna += 'A';
        }
    }
    return dna;

}