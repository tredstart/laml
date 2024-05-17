# LAme Markup Language (laml)

A simple way to write data files that can be compiled to (for now only) zig files. 

# Example 

```laml
struct_name:
    type field,
;

struct_name variable:
    field = value,
;
```

will produce 

```zig 
pub const struct_name = struct{
    field: type,
};
pub var variable = struct_name{
    .field = value,
};
```
