require 'compass/import-once/activate'


Encoding.default_external = Encoding.find('utf-8')
# Require any additional compass plugins here.
# bootstrap 已经包含normalize
# require 'compass-normalize'
# Set this to the root of your project when deployed:
http_path = "/"
# project_path = ""
css_dir = ".tmp/serve/app"
sass_dir = "src/app"
images_dir = "src/assets/images"
javascripts_dir = "src/app/"
sprite_load_path = ["src/assets/images/sprite"]
# You can select your preferred output style here (can be overridden via the command line):
# output_style = :expanded or :nested or :compact or :compressed
# output_style = :expanded
# To enable relative paths to assets via compass helper functions. Uncomment:
# relative_assets = true

# To disable debugging comments that display the original location of your selectors. Uncomment:
# line_comments = false


# If you prefer the indented syntax, you might want to regenerate this
# project again passing --syntax sass, or you can uncomment this:
# preferred_syntax = :sass
# and then run:
# sass-convert -R --from scss --to sass sass scss && rm -rf sass && mv scss sass
sourcemap = true
# 使用绝对路径方便调试http://sass-lang.com/documentation/file.SASS_REFERENCE.html#options
sass_options = {:sourcemap => :file}