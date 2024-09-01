/* Write the body of the getArchitects function, which returns an array containing 2 arrays of HTML elements:

the first array contains the architects, all corresponding to a <a> tag
the second array contains all the non-architects people


Write the body of the getClassical function, which returns an array containing 2 arrays of HTML elements:
the first array contains the architects belonging to the classical class
the second array contains the non-classical architects


Write the body of the getActive function, which returns an array containing 2 arrays of HTML elements:
the first array contains the classical architects who are active in their class
the second array contains the non-active classical architects


Write the body of the getBonannoPisano function, which returns an array containing:
the HTML element of the architect you're looking for, whose id is BonannoPisano
an array which contains all the remaining HTML elements of active classical architects */

export const getArchitects = () =>
  [
    document.querySelectorAll("a")
    , document.querySelectorAll("span")
  ];

export const getClassical = () =>
  [
    document.querySelectorAll("a.classical")
    , document.querySelectorAll("a:not(.classical)")
  ];

export const getActive = () =>
  [
    document.querySelectorAll("a.classical.active")
    , document.querySelectorAll("a.classical:not(.active)")
  ];


export const getBonannoPisano = () =>
  [
    document.querySelectorAll("a#BonannoPisano.classical.active")[0]
    , document.querySelectorAll("a:not(#BonannoPisano).classical.active")
  ];