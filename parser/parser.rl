package parser

import "fmt"

%%{
  machine functional;
  write data;
}%%

func Parse(data string) error {
  cs, p, pe := 0, 0, len(data)
  
  %%{
    sp = " ";

    name    = alpha+;
    integer = digit+;

    main := name sp "=" sp integer;

    write init;
    write exec;
  }%%

  if cs < functional_first_final {
    return fmt.Errorf("could not parse to byte %d", p)
  }

  return nil
}