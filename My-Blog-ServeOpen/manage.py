#!/usr/bin/env python
"""Django's command-line utility for administrative tasks."""
import os
import sys

import logging
from django.conf import settings


logger = logging.getLogger('django')


def main():
    """Run administrative tasks."""
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'luntan.settings')
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line(sys.argv)


if __name__ == '__main__':
    main()
    # handler = logging.StreamHandler()
    # level = {
    #     '0': logging.WARNING,
    #     '1': logging.INFO,
    #     '2': logging.DEBUG,
    # }.get(settings.VERBOSE, logging.WARNING)
    # handler.setLevel(level)
    # logger.addHandler(handler)
    # try:
    #     from django.core.management import execute_from_command_line
    # except ImportError as exc:
    #     raise ImportError(
    #         "Couldn't import Django. Are you sure it's installed and "
    #         "available on your PYTHONPATH environment variable? Did you "
    #         "forget to activate a virtual environment?"
    #     ) from exc
    # execute_from_command_line()




