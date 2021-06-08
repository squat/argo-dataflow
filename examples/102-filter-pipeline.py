from dsls.python import kafka, pipeline

if __name__ == '__main__':
    (pipeline("102-filter")
     .describe("""This is an example of built-in filtering.

Filters are written using expression syntax and must return a boolean.

They have a single variable, `msg`, which is a byte array.

[Learn about expressions](../docs/EXPRESSIONS.md)""")
     .step(
        kafka('input-topic')
            .filter('main', 'string(msg) contains "capybara"')
            .kafka('output-topic')
    )
     .save())
